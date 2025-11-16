import { useState, useEffect } from 'react'
import './App.css'
import UserForm from './UserForm'
import LoginPage from './LoginPage'
import { isLoggedIn, getCurrentUserId, logout } from 'shared'

function App() {
  const [loggedIn, setLoggedIn] = useState(false)
  const [currentUserId, setCurrentUserId] = useState<string | null>(null)

  useEffect(() => {
    // 初期状態を確認
    const loggedInState = isLoggedIn()
    setLoggedIn(loggedInState)
    if (loggedInState) {
      setCurrentUserId(getCurrentUserId())
    }
  }, [])

  const handleLogin = () => {
    setLoggedIn(true)
    setCurrentUserId(getCurrentUserId())
  }

  const handleLogout = () => {
    logout()
    setLoggedIn(false)
    setCurrentUserId(null)
  }

  if (!loggedIn) {
    return <LoginPage onLogin={handleLogin} />
  }

  return (
    <>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h1>Tollo Q&A Platform</h1>
        <div>
          <span style={{ marginRight: '1rem' }}>Logged in as: {currentUserId}</span>
          <button onClick={handleLogout}>Logout</button>
        </div>
      </div>
      <div className="card">
        <UserForm />
      </div>
    </>
  )
}

export default App
