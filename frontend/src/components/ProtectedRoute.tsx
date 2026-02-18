import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { UserRole } from '../types';

interface ProtectedRouteProps {
  children: React.ReactNode;
  role: UserRole;
}

export default function ProtectedRoute({ children, role }: ProtectedRouteProps) {
  const { user, isAuthenticated } = useAuth();

  if (!isAuthenticated) {
    return <Navigate to="/" replace />;
  }

  if (user?.role !== role) {
    return <Navigate to="/" replace />;
  }

  return <>{children}</>;
}
