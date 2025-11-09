import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { greet } from 'shared'; // ここを修正

function App() {
  const [count, setCount] = useState(0)
  const [kotlinGreeting, setKotlinGreeting] = useState(''); // ここを追加

  // コンポーネントがマウントされたときにKotlinの関数を呼び出す
  useState(() => { // useEffect の代わりに useState の初期化で呼び出し
    setKotlinGreeting(greet()); // ここを修正
  });

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
      {/* Kotlinからの挨拶を表示 */}
      <p>{kotlinGreeting}</p> {/* ここを追加 */}
    </>
  )
}

export default App
