import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import api from '../../api/axios';
import { PatientDetails } from '../../types';
import LoadingSpinner from '../../components/LoadingSpinner';
import toast from 'react-hot-toast';
import { User, Star } from 'lucide-react';

export default function PatientProfile() {
  const [profile, setProfile] = useState<PatientDetails>({
    fullname: '', email: '', gender: '', contactnumber: '',
  });
  const [editing, setEditing] = useState<Record<string, boolean>>({});
  const [loading, setLoading] = useState(true);
  const [saving, setSaving] = useState(false);
  const [bookedDoctors, setBookedDoctors] = useState<any[]>([]);

  useEffect(() => {
    api.get('/patient/profile')
      .then((res) => { if (res.data.data) setProfile(res.data.data); })
      .catch(() => toast.error('Failed to load profile'))
      .finally(() => setLoading(false));
    api.get('/patient/booking')
      .then((res) => setBookedDoctors(res.data.data || []))
      .catch(() => {});
  }, []);

  const handleSave = async (field: string) => {
    setSaving(true);
    try {
      const res = await api.put('/patient/profile', profile);
      if (res.data.data) setProfile(res.data.data);
      toast.success('Updated!');
      setEditing({ ...editing, [field]: false });
    } catch (err: any) {
      toast.error(err.response?.data?.error || 'Failed to update');
    } finally {
      setSaving(false);
    }
  };

  if (loading) return <LoadingSpinner />;

  const fields = [
    { key: 'fullname', label: 'full name' },
    { key: 'email', label: 'email' },
    { key: 'contactnumber', label: 'phoneNumber' },
    { key: 'gender', label: 'gender' },
  ];

  return (
    <div className="space-y-8">
      {/* Header */}
      <div className="page-header">
        <h1 className="page-title text-4xl">about me</h1>
        <span className="brand-name text-2xl">LifeLink</span>
      </div>

      {/* Profile card - yellow */}
      <div className="card-yellow text-center">
        <div className="w-20 h-20 rounded-full bg-white border-2 border-gray-200 flex items-center justify-center mx-auto mb-2">
          <User className="h-10 w-10 text-gray-400" />
        </div>
        <p className="font-handwritten text-lg italic">{profile.fullname || 'patient-name'}</p>

        <div className="text-left mt-6 space-y-3">
          {fields.map(({ key, label }) => (
            <div key={key} className="flex items-center justify-between gap-4">
              <div className="flex-1">
                <span className="font-handwritten font-bold text-base">{label}:</span>
                {editing[key] ? (
                  <input
                    type="text"
                    className="ml-2 px-2 py-1 bg-white rounded-lg border-none outline-none text-sm font-sans"
                    value={(profile as any)[key]}
                    onChange={(e) => setProfile({ ...profile, [key]: e.target.value })}
                    onBlur={() => handleSave(key)}
                    autoFocus
                  />
                ) : (
                  <span className="ml-2 text-sm font-sans">{(profile as any)[key] || ''}</span>
                )}
              </div>
              <button
                onClick={() => setEditing({ ...editing, [key]: !editing[key] })}
                className="btn-dark text-xs py-1 px-4"
              >
                {editing[key] ? 'save' : 'edit'}
              </button>
            </div>
          ))}
        </div>
      </div>

      {/* Booked Doctors - gray cards */}
      {bookedDoctors.length > 0 && (
        <div className="space-y-4">
          {bookedDoctors.map((b: any, i: number) => (
            <div key={i} className="card-gray flex items-center gap-4">
              <div className="w-14 h-14 rounded-full bg-white border-2 border-gray-200 flex items-center justify-center flex-shrink-0">
                <User className="h-7 w-7 text-gray-400" />
              </div>
              <div className="flex-1">
                <p className="font-handwritten font-bold text-base">give review:</p>
                <div className="flex gap-0.5 my-1">
                  {[1, 2, 3, 4, 5].map((s) => (
                    <Star key={s} className="h-4 w-4 text-yellow-400" fill="currentColor" />
                  ))}
                </div>
                <p className="text-xs text-gray-500 font-handwritten">scheduled on:</p>
                <p className="font-handwritten italic text-base">{b.DoctorName || 'doctor-name'}</p>
              </div>
              <div className="flex gap-2">
                <Link to="/patient/chat" className="btn-dark text-xs py-1.5 px-4">chat</Link>
                <button className="btn-dark text-xs py-1.5 px-4">online meet</button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
