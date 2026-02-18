import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import api from '../../api/axios';
import { useAuth } from '../../context/AuthContext';
import toast from 'react-hot-toast';

export default function DoctorLogin() {
  const navigate = useNavigate();
  const { login } = useAuth();
  const [loading, setLoading] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      const res = await api.post('/doctor/login', { email, password });
      const data = res.data.data;
      login({
        id: String(data.DoctorDetail.Id),
        name: data.DoctorDetail.FullName,
        email: data.DoctorDetail.Email,
        role: 'doctor',
        accessToken: data.AccessToken,
        refreshToken: data.RefreshToken,
      });
      toast.success('Logged in successfully');
      navigate('/doctor/dashboard');
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Login failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-[85vh] flex flex-col items-center justify-center px-4">
      {/* Header */}
      <div className="w-full max-w-2xl page-header mb-6">
        <h1 className="page-title text-4xl">Doctor's Login</h1>
        <span className="brand-name text-4xl">LifeLink</span>
      </div>

      {/* Yellow Card */}
      <form onSubmit={handleSubmit} className="w-full max-w-2xl card-yellow py-10 px-8 sm:px-12">
        <div className="space-y-6 max-w-lg">
          <div>
            <label className="form-label">email</label>
            <input type="email" className="input-field"
              value={email} onChange={(e) => setEmail(e.target.value)}
              placeholder="Enter your email" required />
          </div>

          <div>
            <label className="form-label">PASSWORD</label>
            <input type="password" className="input-field"
              value={password} onChange={(e) => setPassword(e.target.value)}
              placeholder="Enter your password" required minLength={6} />
          </div>

          <div className="pt-4">
            <button type="submit" disabled={loading} className="btn-blue">
              {loading ? 'Signing in...' : 'SUBMIT'}
            </button>
          </div>
        </div>
      </form>

      {/* Bottom link */}
      <div className="mt-6 text-center font-handwritten text-base text-gray-500">
        <p>
          Don't have an account?{' '}
          <Link to="/doctor/signup" className="text-brand-blue font-bold hover:underline">Sign up</Link>
        </p>
      </div>
    </div>
  );
}
