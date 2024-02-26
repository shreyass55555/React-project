import React, { useState } from 'react';

function UserStats() {
  const [userId, setUserId] = useState('');
  const [stats, setStats] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  const fetchUserStats = async () => {
    setIsLoading(true);
    setError(null);
    setStats(null); // Clear existing stats when fetching new data

    try {
      const response = await fetch(`http://localhost:3000/userStats/${userId}`);
      if (!response.ok) {
        throw new Error('Failed to fetch user stats');
      }
      const data = await response.json();
      setStats(data);
    } catch (error) {
      setError(error.message);
    } finally {
      setIsLoading(false);
    }
  };

  const handleChange = (e) => {
    setUserId(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    fetchUserStats();
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label htmlFor="userId">User ID:</label>
        <input
          type="text"
          id="userId"
          value={userId}
          onChange={handleChange}
          placeholder="Enter User ID"
          required
        />
        <button type="submit" disabled={isLoading}>Get User Stats</button>
      </form>
      {isLoading && <p>Loading...</p>}
      {error && <p>Failed to fetch user stats</p>}
      {stats && (
        <div>
          <h2>User Stats</h2>
          <p>
            <strong>Product ID:</strong> {stats.product_id}
          </p>
          <p>
            <strong>Time Taken:</strong> {stats.time_taken} milliseconds
          </p>
        </div>
      )}
    </div>
  );
}

export default UserStats;

