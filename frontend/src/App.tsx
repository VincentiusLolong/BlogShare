import React, { useEffect, useState } from "react";
import "./App.css";
import Home from "./pages/home";
import Login from "./pages/login";
import Nav from "./components/nav";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Register from "./pages/register";
import Cookies from "js-cookie";

function App() {

  const [account_created, setAccount_created] = useState<string>('');
  const [username, setUsername] = useState<string>('');
  // const [reNavigate, setNavigate] = useState<boolean>(false);

  useEffect(() => {
    const authCookie = Cookies.get('jwt');
    if (authCookie === '') {
      // setNavigate(false);
      console.log("Belum Login");
    } else {
      (async () => {
        const response = await fetch(
          "http://localhost:8080/auth/secure/Homepage",
          {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
            // redirect: "manual",
          }
        );
        if (response.status === 200) {
          const conten = await response.json();
          console.log(conten);
          setAccount_created(conten.data.Result.account_created);
          setUsername(conten.data.Result.username);
          // setNavigate(true);
          console.log("JADIIIII");
        } else {
          // setNavigate(false);
          console.log("NGAK JADIIII");
        }
      })();
    }
  }, []);
  
  return (
    <div className="App">
      <BrowserRouter>
      <Nav name={username} setName={setUsername}/>
        <Routes>
          <Route path="/" Component={() => <Home  name={username} created={account_created}/>}/>
          <Route path="/login" Component={() => <Login name={username} setName={setUsername}/>} />
          <Route path="/register" Component={Register} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
