import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'

import { createBrowserRouter,  RouterProvider } from 'react-router-dom'

import SetupPage from './views/Setup.tsx'
import LoginPage from './views/Login.tsx'
import HomePage from './views/Home.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomePage />,
  },
  {
    path: '/setup',
    element: <SetupPage />,
  },
  {
    path: '/login',
    element: <LoginPage />,
  },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
