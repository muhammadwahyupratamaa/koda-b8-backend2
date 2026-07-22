import { createBrowserRouter, Navigate, RouterProvider } from "react-router-dom";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Users from "./pages/Users";


const router = createBrowserRouter ([
  {
    path:"/",
    element: <Navigate to="/login" />,
  },
  {
    path:"/login",
    element:<Login />,
  },
  {
    path:"register",
    element:<Register />,
  },
  {
    path:"users",
    element: <Users />,
  },
])

function App() {
  return <RouterProvider router={router} />
}

export default App