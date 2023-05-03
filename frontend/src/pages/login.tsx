import React, { SyntheticEvent, useState } from "react";
import { Navigate } from "react-router-dom";

const Login = (props: {name:string, setName: (name: string) => void }) => {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  // const [reNavigate, setNavigate] = useState(false);
  // const redirect = () => {
  //   return <Navigate to="/" />;
  // };

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/auth/sign-in", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        email,
        password,
      }),
    });
    if (response.status === 200) {
      const content = await response.json();
      let user: string = content.data.Result;
      // setNavigate(true)
      props.setName(user);
      // redirect();
    } else {
      props.setName("");
    }
    // console.log({
    //   email:email,
    //   password:password
    // }
  };

  if (props.name.length !== 0){
    return <Navigate to="/" />; 
  } else {
    return (
      <div>
        <form className="form-signin" onSubmit={submit}>
          <img
            className="mb-4"
            src="https://upload.wikimedia.org/wikipedia/commons/thumb/7/7f/Saturn.svg/1110px-Saturn.svg.png"
            alt=""
            width="92"
            height="72"
          />
          <h1 className="h3 mb-3 font-weight-normal">sign in</h1>
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
            Sign in
          </button>
          <p className="mt-5 mb-3 text-muted">&copy; 2023</p>
        </form>
      </div>
    );
  }
  
};

export default Login;
