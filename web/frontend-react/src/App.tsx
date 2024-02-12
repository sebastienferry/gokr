import './App.css'
import { createBrowserRouter, RouterProvider } from "react-router-dom";

import Menu from './Components/Menu';
import HomePage from './Pages/HomePage';
import AboutPage from './Pages/AboutPage';
import LoginPage from './Pages/LoginPage';
import SingupPage from './Pages/SignupPage';

function App() {
  
  const router = createBrowserRouter([
    { path: "/", element: <HomePage /> },
    { path: "/about", element: <AboutPage /> },
    { path: "/login", element: <LoginPage /> },    
    { path: "/signup", element: <SingupPage /> }
  ]);

  return (
    <div className="max-width">
      <Menu></Menu>
      <div className="container">
        <RouterProvider router={router} />
      </div>
    </div>
  )
}

export default App
