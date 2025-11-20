import { useState, useEffect } from 'react'
import { BrowserRouter, Routes, Route, Link, Navigate } from 'react-router-dom'
import './App.css'
import UserForm from './UserForm'
import LoginPage from './LoginPage'
import { ChatListPage } from './features/chat/pages/ChatListPage'
import { ChatDetailPage } from './features/chat/pages/ChatDetailPage'
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
    <BrowserRouter>
      <div style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '16px 24px', borderBottom: '1px solid #e0e0e0', backgroundColor: '#fff' }}>
          <div style={{ display: 'flex', gap: '24px', alignItems: 'center' }}>
            <h1 style={{ margin: 0 }}>Tollo</h1>
            <nav style={{ display: 'flex', gap: '16px' }}>
              <Link to="/" style={{ textDecoration: 'none', color: '#0084ff' }}>Home</Link>
              <Link to="/chat" style={{ textDecoration: 'none', color: '#0084ff' }}>Chats</Link>
            </nav>
          </div>
          <div>
            <span style={{ marginRight: '1rem' }}>
              {user?.name} ({user?.email})
            </span>
            <button onClick={handleLogout}>Logout</button>
          </div>
        </div>

        <div style={{ flex: 1 }}>
          <Routes>
            <Route path="/" element={
              <div style={{ maxWidth: '800px', margin: '0 auto', padding: '24px' }}>
                <h2>Welcome, {user?.name}!</h2>
                <p>Your dashboard will be here...</p>
              </div>
            } />
            <Route path="/chat" element={<ChatListPage />} />
            <Route path="/chat/:chatId" element={<ChatDetailPage />} />
            <Route path="*" element={<Navigate to="/" replace />} />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  )
}

export default App
