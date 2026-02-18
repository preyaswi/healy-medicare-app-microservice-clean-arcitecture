import { useEffect, useState } from 'react';
import api from '../../api/axios';
import { DoctorDetailsUpdate } from '../../types';
import LoadingSpinner from '../../components/LoadingSpinner';
import toast from 'react-hot-toast';
import { User, Star } from 'lucide-react';
import { Link } from 'react-router-dom';

export default function DoctorProfileEdit() {
  const [profile, setProfile] = useState<DoctorDetailsUpdate>({
    FullName: '', Email: '', PhoneNumber: '',
    Specialization: '', YearsOfExperience: 0, Fees: 0,
  });
  const [editing, setEditing] = useState<Record<string, boolean>>({});
  const [loading, setLoading] = useState(true);
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    api.get('/doctor/profile')
      .then((res) => {
        const d = res.data.data;
        if (d) {
          setProfile({
            FullName: d.FullName || '', Email: d.Email || '',
            PhoneNumber: d.PhoneNumber || '', Specialization: d.Specialization || '',
            YearsOfExperience: d.YearsOfExperience || 0, Fees: d.Fees || 0,
          });
        }
      })
      .catch(() => toast.error('Failed to load profile'))
      .finally(() => setLoading(false));
  }, []);

  const handleSave = async (field: string) => {
    setSaving(true);
    try {
      await api.put('/doctor/profile', profile);
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
    { key: 'FullName', label: 'full name' },
    { key: 'Email', label: 'email' },
    { key: 'PhoneNumber', label: 'phoneNumber' },
    { key: 'YearsOfExperience', label: 'Years of experience' },
    { key: 'Specialization', label: 'Specialization' },
    { key: 'Fees', label: 'for 1 hour', suffix: ' rupee' },
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
        <p className="font-handwritten text-lg italic">{profile.FullName || 'doctor-name'}</p>
        <div className="flex justify-center gap-0.5 my-2">
          {[1, 2, 3, 4, 5].map((i) => (
            <Star key={i} className="h-4 w-4 text-yellow-400" fill="currentColor" />
          ))}
        </div>

        <div className="text-left mt-6 space-y-3">
          {fields.map(({ key, label, suffix }) => (
            <div key={key} className="flex items-center justify-between gap-4">
              <div className="flex-1">
                <span className="font-handwritten font-bold text-base">{label}:</span>
                {editing[key] ? (
                  <input
                    type={['YearsOfExperience', 'Fees'].includes(key) ? 'number' : 'text'}
                    className="ml-2 px-2 py-1 bg-white rounded-lg border-none outline-none text-sm font-sans"
                    value={(profile as any)[key]}
                    onChange={(e) => setProfile({ ...profile, [key]: ['YearsOfExperience', 'Fees'].includes(key) ? Number(e.target.value) : e.target.value })}
                    onBlur={() => handleSave(key)}
                    autoFocus
                  />
                ) : (
                  <span className="ml-2 text-sm font-sans">
                    {(profile as any)[key]}{suffix || ''}
                  </span>
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

      {/* Apply leave */}
      <div className="text-center">
        <button className="btn-dark text-lg px-10 py-3">
          apply leave
        </button>
      </div>

      {/* My Patients - gray card */}
      <div>
        <h2 className="font-handwritten text-2xl font-bold mb-4">my patients:</h2>
        <PatientsList />
      </div>
    </div>
  );
}

function PatientsList() {
  const [patients, setPatients] = useState<any[]>([]);

  useEffect(() => {
    api.get('/doctor/patient')
      .then((res) => setPatients(res.data.data || []))
      .catch(() => {});
  }, []);

  if (patients.length === 0) {
    return <div className="card-gray py-6 text-center text-gray-500 font-handwritten text-base">No patients yet</div>;
  }

  return (
    <div className="space-y-4">
      {patients.map((p: any) => (
        <div key={p.BookingId} className="card-gray flex items-center gap-4">
          <div className="w-12 h-12 rounded-full bg-white border-2 border-gray-200 flex items-center justify-center flex-shrink-0">
            <User className="h-6 w-6 text-gray-400" />
          </div>
          <div className="flex-1">
            <p className="font-handwritten italic text-base">{p.Fullname || 'patient-name'}</p>
            <p className="text-xs text-gray-500 font-handwritten">scheduled on:</p>
          </div>
          <div className="flex gap-2">
            <Link to="/doctor/chat" className="btn-dark text-xs py-1.5 px-4">chat</Link>
            <button className="btn-dark text-xs py-1.5 px-4">online meet</button>
          </div>
        </div>
      ))}
    </div>
  );
}
