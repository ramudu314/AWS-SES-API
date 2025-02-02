import React from 'react';

function Stats({ emailsSent }) {
  return (
    <div>
      <h2>Statistics</h2>
      <p>Emails Sent: {emailsSent}</p>
    </div>
  );
}

export default Stats;