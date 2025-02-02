import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './styles.css';

function EmailForm({ onEmailSent }) {
  const [to, setTo] = useState('');
  const [subject, setSubject] = useState('');
  const [body, setBody] = useState('');
  const [message, setMessage] = useState('');
  const [emailLimit, setEmailLimit] = useState(10); 
  const [emailsSent, setEmailsSent] = useState(0); 
  const [emailWarmUp, setEmailWarmUp] = useState(false); 

  // Fetch the backend URL from environment variables
  const backendURL = process.env.REACT_APP_BACKEND_URL || 'https://backend-nlsuoydnv-ramudus-projects.vercel.app/';

  // Fetch email stats on component mount
  useEffect(() => {
    const fetchStats = async () => {
      try {
        const response = await axios.get(`${backendURL}/stats`);
        setEmailsSent(response.data.emails_sent);
        setEmailLimit(response.data.email_limit);
      } catch (error) {
        setMessage('Error fetching statistics: ' + error.message);
      }
    };
    fetchStats();
  }, [backendURL]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post(`${backendURL}/sendEmail`, { 
        to,
        subject,
        body,
      });

      setMessage(response.data.message);
      setTo('');
      setSubject('');
      setBody('');
      setEmailsSent((prev) => prev + 1);
      onEmailSent(); 
    } catch (error) {
      if (error.response) {
        const errorMessage = error.response.data.message || error.response.data.error;
        setMessage(`Error sending email: ${errorMessage}`);
      } else if (error.request) {
        setMessage('Error sending email: No response from server');
      } else {
        setMessage('Error sending email: ' + error.message);
      }
    }
  };

  return (
    <div>
      <h2>Send Email</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>To:</label>
          <input
            type="email"
            value={to}
            onChange={(e) => setTo(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Subject:</label>
          <input
            type="text"
            value={subject}
            onChange={(e) => setSubject(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Body:</label>
          <textarea
            value={body}
            onChange={(e) => setBody(e.target.value)}
            required
          />
        </div>
        <button type="submit" disabled={emailsSent >= emailLimit}>
          Send Email
        </button>
      </form>

      <div>
        <h3>Statistics</h3>
        <p>Emails Sent: {emailsSent}</p>
        <p>Email Limit: {emailLimit}</p>
      </div>

      {message && <p>{message}</p>}
    </div>
  );
}

export default EmailForm;
