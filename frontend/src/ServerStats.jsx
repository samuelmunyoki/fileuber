import React, { useEffect } from "react";
import { StartLocalServer } from "../wailsjs/go/server/LocalServer";
import { useState } from "react";
import { EventsOn } from "../wailsjs/runtime/runtime";
import PieChart from "./components/Chart";

export default function ServerStats() {
  function formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return "0 Bytes";
    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + " MB/s";
  }
  useEffect(() => {
    // Run StartLocalServer only once when the component mounts
    StartLocalServer();
    return () => {
    };
  }, []);

  const [cpudata, setCpudata] = useState();
  const [downloads, setDownloads] = useState(null);

  EventsOn("cpudata", (data) => {
    setCpudata(data)
  });
  EventsOn("filedownloads", (data) => {
    setDownloads(data)
  });
  
  return (
    <div>
      <div className="title-sm"> FileUber ServerStats</div>
      <div className="main-section">
        <div className="options-pane">
          <div>
            <div className="options-pane-item">
              <p>
                Server status: <span className="green">Good</span>
              </p>
            </div>
            <div className="options-pane-item">
              <p>
                CPU Model:{" "}
                <span className="green">
                  {cpudata != null
                    ? cpudata["model"].split(" ")[0] +
                      " " +
                      cpudata["model"].split(" ")[2].split("-")[0]
                    : "n/a"}
                </span>
              </p>
            </div>
            <div className="options-pane-item">
              <p>
                CPU Clock:
                <span className="green">
                  {cpudata != null ? cpudata["freq"] : "n/a"} Mhz
                </span>
              </p>
            </div>

            <div className="options-pane-item">
              <p>
                Set Username: <span className="green">User23</span>
              </p>
            </div>
            <div className="options-pane-item">
              <p>
                Set Password: <span className="green">Passw3rd</span>
              </p>
            </div>
          </div>
          <PieChart></PieChart>

          <button className="btn btn-danger">Shutdown</button>
        </div>
        <div className="line"></div>
        <div className="details">
          {/* <div className="chart">
            <PieChart></PieChart>
          </div> */}
          <div className="table">
            {downloads !== null && downloads.length > 0 ? (
              <table className="downloads-table">
                <thead>
                  <tr>
                    <th>File</th>
                    <th>Speed</th>
                  </tr>
                </thead>
                <tbody>
                  {downloads.map((item, index) => (
                    <tr key={index}>
                      <td>{item.file}</td>
                      <td>{formatBytes(item.speed)}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            ) : (
              <p>No current downloads</p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
