import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import api from '../../api/axios';
import toast from 'react-hot-toast';

export default function PatientSignup() {
  const { user, isAuthenticated } = useAuth();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);
  const [form, setForm] = useState({
    fullname: '', email: '', gender: '', contactnumber: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleReset = () => {
    setForm({ fullname: '', email: '', gender: '', contactnumber: '' });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!isAuthenticated) {
      toast.error('Please sign in with Google first');
      window.location.href = '/api/patient/login';
      return;
    }
    if (!form.fullname || !form.gender || !form.contactnumber) {
      toast.error('Please fill all required fields');
      return;
    }
    setLoading(true);
    try {
      const res = await api.put('/patient/profile', {
        fullname: form.fullname,
        email: form.email || user?.email || '',
        gender: form.gender,
        contactnumber: form.contactnumber,
      });
      if (res.data.data) {
        toast.success('Profile updated successfully!');
        navigate('/patient/dashboard');
      }
    } catch (err: any) {
      toast.error(err.response?.data?.error || err.response?.data?.message || 'Failed to update profile');
    } finally {
      setLoading(false);
    }
  };

  const displayEmail = form.email || user?.email || '';
  const displayName = form.fullname || user?.name || '';

  return (
    <div className="min-h-[85vh] flex flex-col items-center justify-start py-8 px-4">
      {/* Header */}
      <div className="w-full max-w-lg page-header mb-6">
        <h1 className="page-title text-3xl">Patient's signup</h1>
        <span className="brand-name text-3xl">LifeLink</span>
      </div>

      {/* Yellow Card Form */}
      <form onSubmit={handleSubmit} className="w-full max-w-lg card-yellow py-8 px-6 sm:px-10 space-y-5">
        <div>
          <label className="form-label">Full Name</label>
          <input type="text" name="fullname" className="input-field"
            value={displayName} onChange={(e) => setForm({ ...form, fullname: e.target.value })}
            placeholder="Enter your full name" />
        </div>

        <div>
          <label className="form-label">Email</label>
          <input type="email" name="email" className="input-field"
            value={displayEmail} onChange={(e) => setForm({ ...form, email: e.target.value })}
            placeholder="Email from Google account" readOnly={isAuthenticated} />
        </div>

        <div>
          <label className="form-label">Gender</label>
          <select name="gender" className="input-field" value={form.gender} onChange={handleChange}>
            <option value="">Select gender</option>
            <option value="male">Male</option>
            <option value="female">Female</option>
            <option value="other">Other</option>
          </select>
        </div>

        <div>
          <label className="form-label">Contact Number</label>
          <input type="tel" name="contactnumber" className="input-field"
            value={form.contactnumber} onChange={handleChange} placeholder="+1234567890" />
        </div>

        <div className="flex items-center gap-3 pt-4">
          <button type="submit" disabled={loading} className="btn-blue">
            {loading ? 'Saving...' : 'SUBMIT'}
          </button>
          <button type="button" onClick={handleReset} className="btn-blue">
            RESET
          </button>
        </div>
      </form>

      {/* Bottom links */}
      <div className="mt-6 text-center space-y-2">
        <p className="font-handwritten text-base text-gray-500">
          Already have an account?{' '}
          <Link to="/patient/login" className="text-brand-blue font-bold hover:underline">Login here</Link>
        </p>
        <p className="text-xs text-gray-400 font-sans">
          By clicking the Sign Up button you agree to our{' '}
          <span className="text-brand-blue cursor-pointer">terms and condition</span>{' '}
          and <span className="text-brand-blue cursor-pointer">Policy Privacy</span>
        </p>
      </div>
    </div>
  );
}
