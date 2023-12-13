import React from 'react'
import {createRoot} from 'react-dom/client'
import { HashRouter, Routes, Route } from "react-router-dom";
import './style.css'
import App from './App'
import ServerStats from './ServerStats';
import ErrorPage from './ErrorPage';

const container = document.getElementById('root')

const root = createRoot(container)

root.render(
  <HashRouter basename={"/"}>
    <Routes>
      <Route path="/" element={<App />} exact />
      <Route path="/serverstats" element={<ServerStats />} />
      <Route path="/error" element={<ErrorPage />} />
    </Routes>
  </HashRouter>,
  root
);
