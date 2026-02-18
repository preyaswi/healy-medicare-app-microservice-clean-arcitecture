import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import api from '../../api/axios';
import { useAuth } from '../../context/AuthContext';
import toast from 'react-hot-toast';

export default function AdminSignup() {
  const navigate = useNavigate();
  const { login } = useAuth();
  const [loading, setLoading] = useState(false);
  const [form, setForm] = useState({
    firstname: '', lastname: '', email: '', password: '',
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      const res = await api.post('/admin/signup', form);
      const data = res.data.data;
      login({
        id: String(data.Admin.id),
        name: `${data.Admin.firstname} ${data.Admin.lastname}`,
        email: data.Admin.Email,
        role: 'admin',
        accessToken: data.Token,
      });
      toast.success('Account created!');
      navigate('/admin/dashboard');
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Signup failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-[80vh] flex flex-col items-center justify-center px-4">
      <div className="w-full max-w-lg page-header mb-6">
        <h1 className="page-title text-3xl tracking-widest">ADMIN'S SIGNUP</h1>
        <span className="brand-name text-3xl">LifeLink</span>
      </div>

      <form onSubmit={handleSubmit} className="w-full max-w-lg card-yellow py-10 px-8 space-y-5">
        <div className="grid sm:grid-cols-2 gap-4">
          <div>
            <label className="form-label">first name</label>
            <input type="text" className="input-field"
              value={form.firstname} onChange={(e) => setForm({ ...form, firstname: e.target.value })} required />
          </div>
          <div>
            <label className="form-label">last name</label>
            <input type="text" className="input-field"
              value={form.lastname} onChange={(e) => setForm({ ...form, lastname: e.target.value })} required />
          </div>
        </div>
        <div>
          <label className="form-label">email</label>
          <input type="email" className="input-field"
            value={form.email} onChange={(e) => setForm({ ...form, email: e.target.value })} required />
        </div>
        <div>
          <label className="form-label">PASSWORD</label>
          <input type="password" className="input-field"
            value={form.password} onChange={(e) => setForm({ ...form, password: e.target.value })} required minLength={6} />
        </div>
        <button type="submit" disabled={loading} className="btn-blue">
          {loading ? 'creating...' : 'SUBMIT'}
        </button>
        <p className="text-center font-handwritten text-base text-gray-600">
          Already have an account?{' '}
          <Link to="/admin/login" className="text-brand-blue-dark font-bold hover:underline">Log in</Link>
        </p>
      </form>
    </div>
  );
}
