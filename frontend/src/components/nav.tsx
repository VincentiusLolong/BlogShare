import React, { useState } from "react";
import { Link, Navigate } from "react-router-dom";

const Nav = (props: {name:string, setName: (name:string) => void}) => {
  const Logout = async () => {
    await fetch("http://localhost:8080/auth/secure/logout", {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });
    props.setName('')
  }
  let menu;

  if(props.name === '') {
    menu = (
      <ul className="right-nav">
      <li>
        <Link to="/login">Login</Link>
      </li>
      <li>
        <Link to="/register">Register</Link>
      </li>
    </ul>
    );
  } else {
    menu = (
      <ul className="right-nav">
      <li>
        <Link to="/login" onClick={Logout}>Logout</Link>
      </li>
    </ul>
    )
  }
  return (
  //   <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
  //     <Link to='/homepage' className="navbar-brand">
  //       BlogDev
  //     </Link>
  //     <button
  //       className="navbar-toggler"
  //       type="button"
  //       data-toggle="collapse"
  //       data-target="#navbarCollapse"
  //       aria-controls="navbarCollapse"
  //       aria-expanded="false"
  //       aria-label="Toggle navigation"
  //     >
  //       <span className="navbar-toggler-icon"></span>
  //     </button>
  //     <div className="collapse navbar-collapse" id="navbarCollapse">
  //   <ul className="navbar-nav mr-auto">
  //     <li className="nav-item active">
  //       <Link to="/" className="nav-link">
  //         Home
  //       </Link>
  //     </li>
  //     <li className="nav-item">
  //       <Link to="/login" className="nav-link">
  //         Sign-In
  //       </Link>
  //     </li>
  //     <li className="nav-item">
  //       <Link to="/register" className="nav-link">
  //         Sign-Up
  //       </Link>
  //     </li>
  //   </ul>
  // </div>
  //   </nav>

        <nav>
          <ul className="left-nav">
            <li>
              <Link to="/">Home</Link>
            </li>
          </ul>
          {menu}
        </nav>
  )
};

export default Nav;
