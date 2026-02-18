import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { AuthUser, UserRole } from '../types';

interface AuthContextType {
  user: AuthUser | null;
  login: (user: AuthUser) => void;
  logout: () => void;
  isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<AuthUser | null>(null);

  useEffect(() => {
    const stored = localStorage.getItem('healy_user');
    if (stored) {
      setUser(JSON.parse(stored));
    }
  }, []);

  const login = (userData: AuthUser) => {
    setUser(userData);
    localStorage.setItem('healy_user', JSON.stringify(userData));
  };

  const logout = () => {
    setUser(null);
    localStorage.removeItem('healy_user');
  };

  return (
    <AuthContext.Provider value={{ user, login, logout, isAuthenticated: !!user }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}

export function useRequireRole(role: UserRole) {
  const { user } = useAuth();
  return user?.role === role;
}
