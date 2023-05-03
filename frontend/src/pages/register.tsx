import React from 'react'

const Register = () => {
    ;<div>
        <form className="form-signin">
            <img
                className="mb-4"
                src="https://icons.veryicon.com/png/o/miscellaneous/esgcc-basic-icon-library/register-14.png"
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
            />
            <button className="btn btn-lg btn-primary btn-block" type="submit">
                Sign Up
            </button>
            <p className="mt-5 mb-3 text-muted">&copy; 2023</p>
        </form>
    </div>
}

export default Register
