import { useState } from "react";
import { Link } from "react-router-dom";
import { FaLock, FaUserAlt } from "react-icons/fa";
import api from "../services/api";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (email.trim() === "") {
      alert("Email wajib diisi!");
      return;
    }

    if (password.trim() === "") {
      alert("Password wajib diisi!");
      return;
    }

    const formData = new URLSearchParams();

    formData.append("email", email);
    formData.append("password", password);

    try {
      const response = await api.post("/login", formData, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      });
      console.log(response.data);
      alert(response.data.message);
      localStorage.setItem("token", response.data.result.token);
      window.location.href = "/users";
    } catch (error) {
      console.log("Status:", error.response?.status);
      console.log("Data:", error.response?.data);
    }
  };

  return (
    <div className="min-h-screen bg-gray-600 flex items-center justify-center px-6">
      <div className="w-full max-w-md">
        <div className="mb-12">
          <h1 className="text-5xl flex justify-center font-bold text-white">
            Login
          </h1>

          <p className="text-gray-300 flex justify-center mt-4">
            Welcome back, please login to continue.
          </p>
        </div>
        <form onSubmit={handleSubmit} className="space-y-10">
          <div>
            <div className="flex items-center gap-4 border-b border-gray-900 pb-3">
              <FaUserAlt className="text-gray-300" />

              <input
                type="email"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full bg-transparent outline-none text-white placeholder:text-gray-400"
              />
            </div>
          </div>
          <div>
            <div className="flex items-center gap-4 border-b border-gray-900 pb-3">
              <FaLock className="text-gray-300" />

              <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full bg-transparent outline-none text-white placeholder:text-gray-400"
              />
            </div>
          </div>

          <div className="flex items-center gap-3 text-gray-300">
            <input
              type="checkbox"
              id="remember"
              className="accent-orange-500"
            />

            <label htmlFor="remember">Remember me</label>
          </div>

          <button
            type="submit"
            className="w-full cursor-pointer bg-gray-900 hover:bg-black text-white py-4 rounded-full text-xl font-semibold transition duration-300 shadow-lg"
          >
            Login
          </button>
        </form>
        <div className="text-center flex justify-center gap-2 mt-12">
          <p className="text-gray-300">First time here ?</p>
          <Link to="/register" className="text-orange-400 hover:underline">
            Sign Up
          </Link>
        </div>
      </div>
    </div>
  );
}

export default Login;
