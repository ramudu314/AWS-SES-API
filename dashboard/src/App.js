import React, { useState } from 'react';
import EmailForm from './EmailForm';
import Stats from './Stats';

function App() {
  const [emailsSent, setEmailsSent] = useState(0);

  const handleEmailSent = () => {
    setEmailsSent(prevCount => prevCount + 1);
  };

  return (
    <div style={{ padding: '20px' }}>
      <h1>Mock SES API</h1>
      {/* Pass the onEmailSent function as a prop */}
      <EmailForm onEmailSent={handleEmailSent} />
      {/* <Stats emailsSent={emailsSent} /> */}
    </div>
  );
}

export default App;