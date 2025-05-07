import { useState } from 'react'
import './App.css'
import { ToastContainer, toast } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import { BASE_URL } from './config';

function LoginForm({ username, setUsername, password, setPassword, handleSubmit }) {  return (
    <div style={{textAlign: 'center', width: '100vw', display: 'flex', flexDirection: 'column', justifyContent: 'center', alignItems: 'center', height: '100vh', margin: '0', padding: '0', position: 'fixed', top: '0', left: '0'}}>
      <div style={{marginBottom: '20px'}}>
        <h2>我的文件管理</h2>
      </div>
      <form onSubmit={handleSubmit} style={{display: 'flex', flexDirection: 'column', gap: '15px', width: '300px', margin: '0 auto', transform: 'translateY(-10%)'}}>
        <div style={{display: 'flex', flexDirection: 'column', gap: '5px'}}>
          <label>用户名:</label>
          <input 
            type="text" 
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
            style={{padding: '8px', borderRadius: '4px', border: '1px solid #ccc'}}
          />
        </div>
        <div style={{display: 'flex', flexDirection: 'column', gap: '5px'}}>
          <label>密码:</label>
          <input 
            type="password" 
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            style={{padding: '8px', borderRadius: '4px', border: '1px solid #ccc'}}
          />
        </div>
        <button 
          type="submit" 
          style={{padding: '10px', backgroundColor: '#1890ff', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer'}}
        >
          登录
        </button>
      </form>
    </div>
  )
}

export default function Login() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem('token'))

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      const requestData = {
        url: `${BASE_URL}/user/login`,
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
        timestamp: new Date().toISOString()
      }
      

      const response = await fetch(requestData.url, {
        method: requestData.method,
        headers: requestData.headers,
        body: requestData.body
      })
      
      if (!response.ok) {
        const errorData = await response.json();
        toast.error(errorData.msg || '登录失败');
        return;
      }
      const data = await response.json();
      if (data.status === 1) {
        toast.error(data.msg);
        return;
      }
      toast.success(data.msg);
      localStorage.setItem('token', data.data);
      setIsLoggedIn(true);
      window.location.href = '/';
    } catch (error) {
      console.error('登录错误:', error)
      toast.error(error.message)
    }
  }

  const handleLogout = async () => {
    try {
      const requestData = {
        url: `${BASE_URL}/user/logout`,
        method: 'POST',
        headers: {
          'Authorization': localStorage.getItem('token'),
          'Content-Type': 'application/json'
        },
        timestamp: new Date().toISOString()
      }
      

      const response = await fetch(requestData.url, {
        method: requestData.method,
        headers: requestData.headers
      })
      
      const data = await response.json()
      if (data.status === 1) {
        toast.error(data.msg)
        return
      }
      toast.success(data.msg)
      localStorage.removeItem('token')
      setIsLoggedIn(false)
      window.location.href = '/'
    } catch (error) {
      console.error('登出错误:', error)
      toast.error(error.message)
    }
  }

  return (
    <div className="login-container">
      <div >
        <ToastContainer 
          position="top-right"
          autoClose={3000}
          hideProgressBar={false}
          newestOnTop={false}
          closeOnClick
          rtl={false}
          pauseOnFocusLoss
          draggable
          pauseOnHover
        />
        <LoginForm 
          username={username}
          setUsername={setUsername}
          password={password}
          setPassword={setPassword}
          handleSubmit={handleSubmit}
        />
      </div>
    </div>
  )
}