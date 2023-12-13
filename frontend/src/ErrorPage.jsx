import React from 'react'
import { Link } from 'react-router-dom'


export default function ErrorPage() {
  return (
    <div>
      <div>There was an error reacting our Server</div>
      <Link to="/">Back</Link>
    </div>
  );
}
