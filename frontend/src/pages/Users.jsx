import { useEffect, useState } from "react";
import { FaEdit, FaPlus, FaTrash } from "react-icons/fa";
import api from "../services/api";

function Users() {
  const [users, setUsers] = useState([]);
  const token = localStorage.getItem("token");
  const [showForm, setShowForm] = useState(false);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [isEdit, setIsEdit] = useState(false);
  const [selectedId, setSelectedId] = useState(null);

  console.log("TOKEN:", token);

  const getUsers = async () => {
    try {
      const response = await api.get("/users", {
        headers: {
          Authorization: token,
        },
      });

      console.log(response.data);

      setUsers(response.data);
    } catch (error) {
      console.error(error.response?.data || error.message);
    }
  };

  const saveUsers = async () => {
    try {
      const token = localStorage.getItem("token");
      const data = new URLSearchParams();
      data.append("name", name);
      data.append("email", email);
      data.append("password", password);

      if (isEdit) {
        await api.put(`/users/${selectedId}`, data, {
          headers: {
            Authorization: token,
          },
        });
      } else {
        await api.post("/users", data, {
          headers: {
            Authorization: token,
          },
        });
      }

      setName("");
      setEmail("");
      setPassword("");
      setShowForm(false);
      setIsEdit(false);
      setSelectedId(null);

      getUsers();
    } catch (error) {
      console.log(error.response?.data || error.message);
    }
  };

  useEffect(() => {
    getUsers();
  }, []);

  const deleteUser = async (id) => {
    const confirmDelete = window.confirm("Delete this user?");

    if (!confirmDelete) return;
    try {
      const token = localStorage.getItem("token");
      await api.delete(`users/${id}`, {
        headers: {
          Authorization: token,
        },
      });
      getUsers();
    } catch (error) {
      console.log(err);
    }
  };

  const handleEdit = (user) => {
    setIsEdit(true)
    setShowForm(true)
    setSelectedId(user.id)
    setName(user.name)
    setEmail(user.email)
    setPassword("")
  }

  return (
    <div className="min-h-screen px-8 py-10">
      <div className="max-w-4xl mx-auto">
        <div className="flex items-center justify-between mb-10">
          <div>
            <h1 className="text-4xl font-bold text-gray-800">Users</h1>

            <p className="text-gray-500 mt-2">Manage your registered users.</p>
          </div>

          <button
            onClick={() => setShowForm(!showForm)}
            className="flex items-center gap-2 bg-gray-900 text-white px-5 py-3 rounded-lg hover:bg-black transition"
          >
            <FaPlus />
            Add User
          </button>
        </div>
        {showForm && (
          <div className="mx-30 my-10">
            <div className="flex flex-col gap-4">
              <input
                type="text"
                name="name"
                value={name}
                id="name"
                onChange={(e) => setName(e.target.value)}
                placeholder="Masukan Nama anda..."
                className="w-full border rounded-xl p-3"
              />

              <input
                type="email"
                name="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                id="email"
                placeholder="Masukan Email anda..."
                className="w-full border rounded-xl p-3"
              />

              <input
                type="password"
                name="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                id="password"
                placeholder="Masukan Password anda..."
                className="w-full border rounded-xl p-3"
              />
            </div>

            <div className=" flex gap-5 text-md mt-5 justify-around">
              <button
                className="border-b border-orange-500"
                onClick={saveUsers}
              >
                {isEdit ? "Update user" : "Save User"}
              </button>

              <button
                className="border-b border-orange-500"
                onClick={() => setShowForm(false)}
              >
                Cancel
              </button>
            </div>
          </div>
        )}

        <div className="space-y-5">
          {users.length === 0 && (
            <p className="text-center text-gray-500">No user data.</p>
          )}
          {users.map((user) => (
            <div
              key={user.id}
              className="flex items-center justify-between border-b border-gray-300 pb-5"
            >
              <div>
                <h2 className="font-semibold text-xl">{user.Name}</h2>

                <p className="text-gray-500">{user.Email}</p>
              </div>

              <div className="flex gap-3">
                <button onClick={() => handleEdit(user)} className="flex items-center gap-2 text-blue-600 hover:text-blue-800">
                  <FaEdit />
                  Edit
                </button>

                <button
                  onClick={() => deleteUser(user.id)}
                  className="flex items-center gap-2 text-red-600 hover:text-red-800"
                >
                  <FaTrash />
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default Users;
