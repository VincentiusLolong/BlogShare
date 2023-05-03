import React from 'react'
import logo from './logo.svg'
import './App.css'
import Login from './pages/login'
import Nav from './components/nav'

function App() {
    return (
        <div className="App">
            <Nav />
            <Login />
        </div>
    )
}

export default App
