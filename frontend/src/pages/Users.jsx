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
  const [picture, setPicture] = useState(null);

  const [isEdit, setIsEdit] = useState(false);
  const [selectedId, setSelectedId] = useState(null);

  const getUsers = async () => {
    try {
      const response = await api.get("/users", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      setUsers(response.data);
    } catch (error) {
      console.log(error.response?.data || error.message);
    }
  };

  useEffect(() => {
    getUsers();
  }, []);

  const saveUsers = async () => {
    try {
      const data = new FormData();

      data.append("name", name);
      data.append("email", email);
      data.append("password", password);

      if (picture) {
        data.append("picture", picture);
      }

      if (isEdit) {
        await api.put(`/users/${selectedId}`, data, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
      } else {
        await api.post("/users", data, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
      }

      setName("");
      setEmail("");
      setPassword("");
      setPicture(null);

      setShowForm(false);
      setIsEdit(false);
      setSelectedId(null);

      getUsers();
    } catch (error) {
      console.log(error.response?.data || error.message);
    }
  };

  const deleteUser = async (id) => {
    const confirmDelete = window.confirm("Delete this user?");

    if (!confirmDelete) return;

    try {
      await api.delete(`/users/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      getUsers();
    } catch (error) {
      console.log(error.response?.data || error.message);
    }
  };

  const handleEdit = (user) => {
    setIsEdit(true);
    setShowForm(true);

    setSelectedId(user.id);

    setName(user.name);
    setEmail(user.email);
    setPassword("");
    setPicture(null);
  };

  return (
    <div className="min-h-screen px-8 py-10">
      <div className="max-w-4xl mx-auto">
        <div className="flex items-center justify-between mb-10">
          <div>
            <h1 className="text-4xl font-bold text-gray-800">Users</h1>

            <p className="text-gray-500 mt-2">Manage your registered users.</p>
          </div>

          <button
            onClick={() => {
              setShowForm(!showForm);

              if (showForm) {
                setName("");
                setEmail("");
                setPassword("");
                setPicture(null);
                setIsEdit(false);
                setSelectedId(null);
              }
            }}
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
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Masukan Nama..."
                className="w-full border rounded-xl p-3"
              />

              <input
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="Masukan Email..."
                className="w-full border rounded-xl p-3"
              />

              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Masukan Password..."
                className="w-full border rounded-xl p-3"
              />

              <input
                type="file"
                accept="image/*"
                onChange={(e) => setPicture(e.target.files[0])}
                className="w-full border rounded-xl p-3"
              />
            </div>

            <div className="flex gap-5 mt-5 justify-around">
              <button
                onClick={saveUsers}
                className="border-b border-orange-500"
              >
                {isEdit ? "Update User" : "Save User"}
              </button>

              <button
                className="border-b border-orange-500"
                onClick={() => {
                  setShowForm(false);

                  setName("");
                  setEmail("");
                  setPassword("");
                  setPicture(null);

                  setIsEdit(false);
                  setSelectedId(null);
                }}
              >
                Cancel
              </button>
            </div>
          </div>
        )}

        <div className="bg-white rounded-xl shadow border overflow-hidden">
          <table className="w-full">
            <thead className="bg-gray-100">
              <tr>
                <th className="text-left px-6 py-4">Photo</th>
                <th className="text-left px-6 py-4">Name</th>
                <th className="text-left px-6 py-4">Email</th>
                <th className="text-center px-6 py-4">Action</th>
              </tr>
            </thead>

            <tbody>
              {users.length === 0 ? (
                <tr>
                  <td colSpan="4" className="text-center py-10 text-gray-500">
                    No user data.
                  </td>
                </tr>
              ) : (
                users.map((user) => (
                  <tr key={user.id} className="border-t hover:bg-gray-50">
                    <td className="px-6 py-4">
                      <img
                        src={
                          user.picture
                            ? `http://localhost:8080/uploads/${user.picture}`
                            : "https://ui-avatars.com/api/?name=User"
                        }
                        alt={user.name}
                        className="w-12 h-12 rounded-full object-cover border"
                      />
                    </td>

                    <td className="px-6 py-4 font-medium">{user.name}</td>

                    <td className="px-6 py-4 text-gray-600">{user.email}</td>

                    <td className="px-6 py-4">
                      <div className="flex justify-center gap-5">
                        <button
                          onClick={() => handleEdit(user)}
                          className="flex items-center gap-2 text-blue-600 hover:text-blue-800"
                        >
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
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}

export default Users;
