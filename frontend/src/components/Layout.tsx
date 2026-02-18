import { Outlet } from 'react-router-dom';
import Navbar from './Navbar';
import Sidebar from './Sidebar';
import { useAuth } from '../context/AuthContext';

export default function Layout() {
  const { isAuthenticated } = useAuth();

  return (
    <div className="min-h-screen bg-white">
      <Navbar />
      <div className="flex">
        {isAuthenticated && <Sidebar />}
        <main className="flex-1 p-6 lg:p-8 max-w-7xl mx-auto w-full">
          <Outlet />
        </main>
      </div>
    </div>
  );
}
