import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { LineChart, Line, CartesianGrid, XAxis, YAxis, Tooltip, Legend, ResponsiveContainer } from 'recharts';

const MarksGraph = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    // Fetch data from your Go API endpoint
    axios.get('http://localhost:8080/api/marks')
      .then((response) => {
        setData(response.data);
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }, []);

  return (
    <div>
      <h2>Marks vs. Name</h2>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={data}>
          <CartesianGrid stroke="#ccc" strokeDasharray="5 5" />
          <XAxis dataKey="name" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Line type="monotone" dataKey="marks" stroke="#8884d8" name="Marks" />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
};

export default MarksGraph;
