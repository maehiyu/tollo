import { useState, useEffect } from 'react'
import './App.css'
import UserForm from './UserForm'
import LoginPage from './LoginPage'
import { isLoggedIn, getCurrentUser, logout, JsUser } from 'shared'

function App() {
  const [loggedIn, setLoggedIn] = useState(false)
  const [user, setUser] = useState<JsUser | null>(null)
  const [needsRegistration, setNeedsRegistration] = useState(false)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const checkAuth = async () => {
      const loggedInState = isLoggedIn()
      setLoggedIn(loggedInState)

      if (loggedInState) {
        // meクエリでユーザー情報を取得
        const currentUser = await getCurrentUser()
        if (currentUser) {
          setUser(currentUser)
        } else {
          // ユーザーが存在しない → 登録が必要
          setNeedsRegistration(true)
        }
      }
      setLoading(false)
    }

    checkAuth()
  }, [])

  const handleLogin = async () => {
    setLoggedIn(true)
    setLoading(true)

    // ログイン後にユーザー情報を取得
    const currentUser = await getCurrentUser()
    if (currentUser) {
      setUser(currentUser)
    } else {
      setNeedsRegistration(true)
    }
    setLoading(false)
  }

  const handleLogout = () => {
    logout()
    setLoggedIn(false)
    setUser(null)
    setNeedsRegistration(false)
  }

  const handleRegistrationComplete = async () => {
    // 登録完了後、ユーザー情報を再取得
    const currentUser = await getCurrentUser()
    if (currentUser) {
      setUser(currentUser)
      setNeedsRegistration(false)
    }
  }

  if (loading) {
    return <div>Loading...</div>
  }

  if (!loggedIn) {
    return <LoginPage onLogin={handleLogin} />
  }

  if (needsRegistration) {
    return (
      <>
        <h1>Welcome! Please complete your registration</h1>
        <UserForm onComplete={handleRegistrationComplete} />
      </>
    )
  }

  return (
    <>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h1>Tollo Q&A Platform</h1>
        <div>
          <span style={{ marginRight: '1rem' }}>
            Logged in as: {user?.name} ({user?.email})
          </span>
          <button onClick={handleLogout}>Logout</button>
        </div>
      </div>
      <div className="card">
        <h2>Welcome, {user?.name}!</h2>
        <p>Your dashboard will be here...</p>
      </div>
    </>
  )
}

export default App
