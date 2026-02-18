import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import api from '../../api/axios';
import { useAuth } from '../../context/AuthContext';
import { DoctorSignUp } from '../../types';
import toast from 'react-hot-toast';

export default function DoctorSignup() {
  const navigate = useNavigate();
  const { login } = useAuth();
  const [loading, setLoading] = useState(false);
  const [form, setForm] = useState<DoctorSignUp>({
    full_name: '', email: '', phone_number: '', password: '',
    confirm_password: '', specialization: '', years_of_experience: 0,
    license_number: '', fees: 0,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: ['years_of_experience', 'fees'].includes(name) ? Number(value) : value,
    }));
  };

  const handleReset = () => {
    setForm({
      full_name: '', email: '', phone_number: '', password: '',
      confirm_password: '', specialization: '', years_of_experience: 0,
      license_number: '', fees: 0,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (form.password !== form.confirm_password) {
      toast.error('Passwords do not match');
      return;
    }
    setLoading(true);
    try {
      const res = await api.post('/doctor/signup', form);
      const data = res.data.data;
      login({
        id: String(data.DoctorDetail.Id),
        name: data.DoctorDetail.FullName,
        email: data.DoctorDetail.Email,
        role: 'doctor',
        accessToken: data.AccessToken,
        refreshToken: data.RefreshToken,
      });
      toast.success('Account created successfully!');
      navigate('/doctor/dashboard');
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Signup failed');
    } finally {
      setLoading(false);
    }
  };

  const fields = [
    { name: 'full_name', label: 'Full Name', type: 'text' },
    { name: 'email', label: 'email', type: 'email' },
    { name: 'phone_number', label: 'Phone Number', type: 'tel' },
    { name: 'password', label: 'Password', type: 'password' },
    { name: 'confirm_password', label: 'Confirm Password', type: 'password' },
    { name: 'specialization', label: 'Specialization', type: 'text' },
    { name: 'years_of_experience', label: 'Years of Experience', type: 'number' },
    { name: 'license_number', label: 'License Number', type: 'text' },
    { name: 'fees', label: 'Consultation Fee', type: 'number' },
  ];

  return (
    <div className="min-h-[85vh] flex flex-col items-center justify-start py-8 px-4">
      {/* Header */}
      <div className="w-full max-w-lg page-header mb-6">
        <h1 className="page-title text-3xl">Doctor's Signup</h1>
        <span className="brand-name text-3xl">LifeLink</span>
      </div>

      {/* Yellow Card */}
      <form onSubmit={handleSubmit} className="w-full max-w-lg card-yellow py-8 px-6 sm:px-10 space-y-5">
        {fields.map(({ name, label, type }) => (
          <div key={name}>
            <label className="form-label">{label}</label>
            <input type={type} name={name} className="input-field"
              value={(form as any)[name]} onChange={handleChange}
              required minLength={type === 'password' ? 6 : undefined}
              min={type === 'number' ? 0 : undefined} />
          </div>
        ))}

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
          <Link to="/doctor/login" className="text-brand-blue font-bold hover:underline">Login here</Link>
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
