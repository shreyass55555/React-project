// App.js

import React, { useState } from 'react';
import axios from 'axios';

function App() {
  const [name, setName] = useState('');
  const [id, setId] = useState('');
  const [salary, setSalary] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post('http://localhost:8080/api/data', { name, id: parseInt(id), salary:parseInt(salary) }); // Convert id to integer
      setMessage('Data inserted successfully!');
    } catch (error) {
      setMessage('Error inserting data: ' + error.response.data.error);
    }
  };

  return (
    <div>
      <h1>Insert Data into ScyllaDB Table</h1>
      <form onSubmit={handleSubmit}>
        <input type="text" value={name} onChange={(e) => setName(e.target.value)} placeholder="Name" required />
        <input type="number" value={id} onChange={(e) => setId(e.target.value)} placeholder="ID" required />
        <input type="number" value={salary} onChange={(e) => setSalary(e.target.value)} placeholder="Salary" required />
        <button type="submit">Submit</button>
      </form>
      <p>{message}</p>
    </div>
  );
}

export default App;
