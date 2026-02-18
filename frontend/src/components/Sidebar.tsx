import { Link, useLocation } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import {
  LayoutDashboard, User, Users, Stethoscope, Calendar,
  MessageSquare, ClipboardList, FileText
} from 'lucide-react';

interface NavItem {
  label: string;
  path: string;
  icon: React.ReactNode;
}

export default function Sidebar() {
  const { user } = useAuth();
  const location = useLocation();

  const getNavItems = (): NavItem[] => {
    switch (user?.role) {
      case 'patient':
        return [
          { label: 'Home', path: '/patient/dashboard', icon: <LayoutDashboard className="h-5 w-5" /> },
          { label: 'About Me', path: '/patient/profile', icon: <User className="h-5 w-5" /> },
          { label: 'Doctors', path: '/patient/doctors', icon: <Stethoscope className="h-5 w-5" /> },
          { label: 'Chat', path: '/patient/chat', icon: <MessageSquare className="h-5 w-5" /> },
        ];
      case 'doctor':
        return [
          { label: 'Home', path: '/doctor/dashboard', icon: <LayoutDashboard className="h-5 w-5" /> },
          { label: 'About Me', path: '/doctor/profile', icon: <User className="h-5 w-5" /> },
          { label: 'Schedule', path: '/doctor/availability', icon: <Calendar className="h-5 w-5" /> },
          { label: 'Patients', path: '/doctor/patients', icon: <ClipboardList className="h-5 w-5" /> },
          { label: 'Prescription', path: '/doctor/prescriptions', icon: <FileText className="h-5 w-5" /> },
          { label: 'Chat', path: '/doctor/chat', icon: <MessageSquare className="h-5 w-5" /> },
        ];
      case 'admin':
        return [
          { label: 'Dashboard', path: '/admin/dashboard', icon: <LayoutDashboard className="h-5 w-5" /> },
          { label: 'Patients', path: '/admin/patients', icon: <Users className="h-5 w-5" /> },
          { label: 'Doctors', path: '/admin/doctors', icon: <Stethoscope className="h-5 w-5" /> },
        ];
      default:
        return [];
    }
  };

  const navItems = getNavItems();

  return (
    <aside className="w-56 min-h-[calc(100vh-4rem)] hidden lg:block pt-4 pl-4">
      <nav className="space-y-1">
        {navItems.map((item) => {
          const isActive = location.pathname === item.path;
          return (
            <Link
              key={item.path}
              to={item.path}
              className={`flex items-center gap-3 px-4 py-2.5 rounded-full font-handwritten text-base transition-all ${
                isActive
                  ? 'bg-brand-yellow text-brand-black font-bold'
                  : 'text-gray-500 hover:bg-brand-yellow-pale hover:text-brand-black'
              }`}
            >
              <span className={isActive ? 'text-brand-black' : 'text-gray-400'}>{item.icon}</span>
              {item.label}
            </Link>
          );
        })}
      </nav>
    </aside>
  );
}
