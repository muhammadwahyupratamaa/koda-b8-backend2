import { useState } from "react";
import { Link } from "react-router-dom";
import { FaUserAlt, FaEnvelope, FaLock } from "react-icons/fa";
import api from "../services/api";

function Register() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (name.trim() === "") {
      alert("Nama wajib diisi!");
      return;
    }

    if (email.trim() === "") {
      alert("Email wajib diisi!");
      return;
    }

    if (password.trim() === "") {
      alert("Password wajib diisi!");
      return;
    }

    if (confirmPassword.trim() === "") {
      alert("Konfirmasi password wajib diisi!");
      return;
    }

    if (password !== confirmPassword) {
      alert("Password tidak sama!");
      return;
    }

    const formData = new URLSearchParams();

    formData.append("name", name);
    formData.append("email", email);
    formData.append("password", password);

    try {
      const response = await api.post("/register", formData, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
      });

      alert(response.data.message);
      localStorage.setItem("token", "hello");
      window.location.href = "/login";
    } catch (error) {
      console.log("Status:", error.response?.status);
      console.log("Data:", error.response?.data);
    }
  };

  return (
    <div className="min-h-screen bg-gray-600 flex items-center justify-center px-6">
      <div className="w-full max-w-md">
        <div className="mb-10 flex flex-col justify-center items-center">
          <h1 className="text-5xl font-bold text-white">Register</h1>

          <p className="text-gray-300 mt-4">Create your new account.</p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-8">
          <div className="flex items-center gap-4 border-b border-gray-900 pb-3">
            <FaUserAlt className="text-gray-300" />

            <input
              type="text"
              placeholder="Full Name"
              className="w-full bg-transparent outline-none text-white placeholder:text-gray-400"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>

          <div className="flex items-center gap-4 border-b border-gray-900 pb-3 ">
            <FaEnvelope className="text-gray-300" />

            <input
              type="email"
              placeholder="Email"
              className="w-full bg-transparent outline-none text-white placeholder:text-gray-400"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>

          <div className="flex items-center gap-4 border-b border-gray-900 pb-3 ">
            <FaLock className="text-gray-300" />

            <input
              type="password"
              placeholder="Password"
              className="w-full bg-transparent outline-none text-white placeholder:text-gray-400"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>

          <div className="flex items-center gap-4 border-b border-gray-900 pb-3 ">
            <FaLock className="text-gray-300" />

            <input
              type="password"
              placeholder="Confirm Password"
              className="w-full bg-transparent outline-none text-white placeholder:text-gray-400"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
            />
          </div>

          <button
            type="submit"
            className="w-full cursor-pointer bg-gray-900 hover:bg-black text-white py-4 rounded-full text-lg font-semibold transition duration-300 shadow-lg"
          >
            Create Account
          </button>
        </form>

        <div className="text-center mt-10 flex justify-center gap-2">
          <p className="text-gray-300">Already have an account ?</p>
          <Link
            to="/login"
            className="text-orange-400 hover:underline cursor-pointer"
          >
            Login
          </Link>
        </div>
      </div>
    </div>
  );
}

export default Register;
