import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { LogOut, User, Menu, X } from 'lucide-react';
import { useState } from 'react';

export default function Navbar() {
  const { user, logout, isAuthenticated } = useAuth();
  const navigate = useNavigate();
  const [mobileOpen, setMobileOpen] = useState(false);

  const handleLogout = () => {
    logout();
    navigate('/');
  };

  const getDashboardLink = () => {
    if (!user) return '/';
    switch (user.role) {
      case 'patient': return '/patient/dashboard';
      case 'doctor': return '/doctor/dashboard';
      case 'admin': return '/admin/dashboard';
      default: return '/';
    }
  };

  return (
    <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-sm">
      <div className="max-w-6xl mx-auto px-4 sm:px-6 py-4">
        <div className="flex justify-between items-center">
          <Link to={getDashboardLink()} className="brand-name text-3xl no-underline">
            LifeLink
          </Link>

          {/* Desktop */}
          <div className="hidden md:flex items-center gap-4">
            {isAuthenticated ? (
              <div className="flex items-center gap-4">
                <Link to={getDashboardLink()} className="p-2 bg-brand-black rounded-full text-white hover:bg-gray-700 transition-colors">
                  <User className="h-5 w-5" />
                </Link>
                <button onClick={handleLogout} className="flex items-center gap-1.5 text-gray-500 hover:text-red-600 transition-colors font-handwritten text-base">
                  <LogOut className="h-4 w-4" />
                </button>
              </div>
            ) : (
              <div className="flex items-center gap-3">
                <Link to="/doctor/login" className="font-handwritten text-base text-gray-600 hover:text-brand-black transition-colors">
                  Doctor
                </Link>
                <Link to="/admin/login" className="font-handwritten text-base text-gray-600 hover:text-brand-black transition-colors">
                  Admin
                </Link>
                <a href="/api/patient/login" className="btn-dark text-sm py-2 px-5">
                  Patient Login
                </a>
              </div>
            )}
          </div>

          {/* Mobile toggle */}
          <div className="md:hidden flex items-center">
            <button onClick={() => setMobileOpen(!mobileOpen)} className="text-gray-600">
              {mobileOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
            </button>
          </div>
        </div>
      </div>

      {/* Mobile menu */}
      {mobileOpen && (
        <div className="md:hidden bg-white px-4 py-3 space-y-3 border-t border-gray-100">
          {isAuthenticated ? (
            <>
              <div className="flex items-center gap-2 font-handwritten text-base">
                <User className="h-4 w-4" />
                <span className="font-bold">{user?.name}</span>
                <span className="text-xs px-2 py-0.5 bg-brand-yellow rounded-full capitalize font-sans">{user?.role}</span>
              </div>
              <button onClick={handleLogout} className="flex items-center gap-1.5 text-red-600 font-handwritten text-base">
                <LogOut className="h-4 w-4" /> Logout
              </button>
            </>
          ) : (
            <>
              <a href="/api/patient/login" className="block font-handwritten text-base font-bold">Patient Login</a>
              <Link to="/doctor/login" className="block font-handwritten text-base text-gray-600">Doctor Login</Link>
              <Link to="/admin/login" className="block font-handwritten text-base text-gray-600">Admin Login</Link>
            </>
          )}
        </div>
      )}
    </nav>
  );
}
