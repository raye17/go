import { useState } from 'react'
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import Login from './Login'
import ImageGallery from './ImageGallery'
import './App.css'
import { BASE_URL } from './config';

export default function App() {
  // 检查token是否为空
  if (!localStorage.getItem('token') && window.location.pathname !== '/login') {
    window.location.href = '/login';
    return null;
  }

  const handleFileUpload = async () => {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = 'image/jpeg,image/png,image/gif,video/mp4,video/avi,video/mov';
    input.onchange = async (e) => {
      const file = e.target.files[0];
      
      const formData = new FormData();
      formData.append('file', file);
      
      try {
        const response = await fetch(`${BASE_URL}/img/upload`, {
          method: 'POST',
          headers: {
            'Authorization': `${localStorage.getItem('token')}`
          },
          body: formData
        });
        const data = await response.json();
        alert(data.msg);
      } catch (error) {
        console.error('上传错误:', error);
        alert('上传失败');
      }
    };
    input.click();
  };

  

  const handleLogout = async () => {
    try {
      const response = await fetch(`${BASE_URL}/user/logout`, {
        method: 'POST',
        headers: {
          'Authorization': localStorage.getItem('token'),
          'Content-Type': 'application/json'
        }
      });
      
      const data = await response.json();
      if (data.status === 1) {
        alert(data.msg);
        return;
      }
      localStorage.removeItem('token');
      window.location.href = '/';
    } catch (error) {
      console.error('登出错误:', error);
      alert('登出失败');
    }
  };

  return (
    <Router>
      <Routes>
        <Route path="/" element={!!localStorage.getItem('token') ? <Navigate to="/resource" replace /> : <Navigate to="/login" replace />} />
        <Route path="/resource" element={
          <div className="App">
            <h1 className="main-title">我的文件管理</h1>
            <button className="upload-button" onClick={handleFileUpload}>上传文件</button>
            <button className="logout-button" onClick={handleLogout}>登出</button>
            <ImageGallery />
          </div>
        } />
        
        <Route path="/login" element={<Login />} />
      </Routes>
    </Router>
  )
}
