import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import api from '../../api/axios';
import { useAuth } from '../../context/AuthContext';
import toast from 'react-hot-toast';

export default function AdminLogin() {
  const navigate = useNavigate();
  const { login } = useAuth();
  const [loading, setLoading] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      const res = await api.post('/admin/login', { email, password });
      const data = res.data.data;
      login({
        id: String(data.Admin.id),
        name: `${data.Admin.firstname} ${data.Admin.lastname}`,
        email: data.Admin.Email,
        role: 'admin',
        accessToken: data.Token,
      });
      toast.success('Logged in successfully');
      navigate('/admin/dashboard');
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Login failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-[80vh] flex flex-col items-center justify-center px-4">
      <div className="w-full max-w-lg page-header mb-6">
        <h1 className="page-title text-3xl tracking-widest">ADMIN'S LOGIN</h1>
        <span className="brand-name text-3xl">LifeLink</span>
      </div>

      <form onSubmit={handleSubmit} className="w-full max-w-lg card-yellow py-10 px-8 space-y-6">
        <div>
          <label className="form-label">email</label>
          <input type="email" className="input-field"
            value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>
        <div>
          <label className="form-label">PASSWORD</label>
          <input type="password" className="input-field"
            value={password} onChange={(e) => setPassword(e.target.value)} required minLength={6} />
        </div>
        <button type="submit" disabled={loading} className="btn-blue">
          {loading ? 'signing in...' : 'SUBMIT'}
        </button>
        <p className="text-center font-handwritten text-base text-gray-600">
          Need an admin account?{' '}
          <Link to="/admin/signup" className="text-brand-blue-dark font-bold hover:underline">Sign up</Link>
        </p>
      </form>
    </div>
  );
}
