
import { useState } from "react";
import "./App.css";
import { useNavigate } from "react-router-dom";
import { SetShareConfig } from "../wailsjs/go/server/ShareConfig";



function App() {
  const [path, setPath] = useState();
  const [username, setUsername] = useState();
  const [password, setPassword] = useState();
  const [port, setPort] = useState();
  const navigate = useNavigate();
  let handleSubmit = async (e) => {
    e.preventDefault();
    
    SetShareConfig(path, username, password, port)
    
    
    setPath("");
    setPassword("");
    setPort("");
    setUsername("");
    navigate("/serverstats");
    
    
  };

  return (
    <div id="App">
      <div className="container">
        <p className="title">FileUber</p>
        <form action="/" method="post" onSubmit={handleSubmit}>
          <div className="input-group">
            <h5 className="input-label">Path</h5>
            <input
              type="text"
              name="path"
              placeholder="path/folder/file"
              required
              value={path}
              onChange={(e) => setPath(e.target.value)}
            />
          </div>
          <div className="input-group">
            <h5 className="input-label">Username</h5>
            <input
              type="text"
              name="username"
              placeholder="fileub3r"
              required
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div className="input-group">
            <h5 className="input-label">Password</h5>
            <input
              type="password"
              name="password"
              placeholder="******"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className="input-group">
            <h5 className="input-label">Port</h5>
            <input
              type="number"
              name="port"
              placeholder="3037"
              required
              value={port}
              onChange={(e) => setPort(e.target.value)}
            />
          </div>

          <div className="input-btn">
            <button type="submit" className="btn active">
              Next
            </button>
          </div>
        </form>
        <a
          href="https://www.google.com"
          target="_blank"
          rel="noopener noreferrer"
          className="terms-link"
        >
          Terms of service
        </a>
      </div>
    </div>
  );
}

export default App;
