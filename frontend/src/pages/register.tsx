import React, { SyntheticEvent, useState } from "react";
import { Navigate } from "react-router-dom";

const Register = () => {
  const [username, setName] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [birth_date, setDate] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [reNavigate, setNavigate] = useState<boolean>(false);

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();
    // console.log({
    //   name: username,
    //   email:email,
    //   date:birth_date,
    //   password:password
    // })
    const response = await fetch("http://localhost:8080/auth/sign-up", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email,
        username,
        password,
        birth_date,
      }),
    });
    setNavigate(true);
    const content = await response.json();
    console.log(content);
  };
  if (reNavigate) {
    return <Navigate to="/login" />;
  } else {
    return (
      <div>
        <form className="form-signin" onSubmit={submit}>
          <img
            className="mb-4"
            src="https://icons.veryicon.com/png/o/miscellaneous/esgcc-basic-icon-library/register-14.png"
            alt=""
            width="64"
            height="64"
          />
          <h1 className="h3 mb-3 font-weight-normal">Sign Up</h1>
          <label htmlFor="inputEmail" className="sr-only">
            Email address
          </label>
          <input
            type="email"
            id="inputEmail"
            className="form-control"
            placeholder="Email address"
            required
            autoFocus
            onChange={(e) => setEmail(e.target.value)}
          />
          <label htmlFor="inputUsername" className="sr-only">
            Username
          </label>
          <input
            type="username"
            id="inputUsername"
            className="form-control"
            placeholder="Username"
            required
            onChange={(e) => setName(e.target.value)}
          />
          <label htmlFor="inputDate" className="sr-only">
            Date
          </label>
          <input
            type="date"
            id="inputDate"
            className="form-control"
            placeholder="Date"
            required
            onChange={(e) => setDate(e.target.value)}
          />
          <label htmlFor="inputPassword" className="sr-only">
            Password
          </label>
          <input
            type="password"
            id="inputPassword"
            className="form-control"
            placeholder="Password"
            required
            onChange={(e) => setPassword(e.target.value)}
          />
          <button className="btn btn-lg btn-primary btn-block" type="submit">
            Sign Up
          </button>
          <p className="mt-5 mb-3 text-muted">&copy; 2023</p>
        </form>
      </div>
    );
  }
};

export default Register;
