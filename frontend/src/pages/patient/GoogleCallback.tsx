import { useEffect, useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { useAuth } from '../../context/AuthContext';
import LoadingSpinner from '../../components/LoadingSpinner';
import api from '../../api/axios';
import toast from 'react-hot-toast';
import { XCircle } from 'lucide-react';

export default function GoogleCallback() {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const { login } = useAuth();
  const [error, setError] = useState('');

  useEffect(() => {
    const handleCallback = async () => {
      const code = searchParams.get('code');
      if (!code) {
        setError('No authorization code received');
        return;
      }

      try {
        const res = await api.get(`/google/redirect?code=${code}`);
        const data = res.data.data;

        login({
          id: data.Patient.Id,
          name: data.Patient.FullName,
          email: data.Patient.Email,
          role: 'patient',
          accessToken: data.AccessToken,
          refreshToken: data.RefreshToken,
        });

        toast.success('Logged in successfully!');

        if (!data.Patient.FullName || data.Patient.FullName === data.Patient.Email) {
          navigate('/patient/signup');
        } else {
          navigate('/patient/dashboard');
        }
      } catch (err: any) {
        setError(err.response?.data?.error || 'Authentication failed');
      }
    };

    handleCallback();
  }, [searchParams, login, navigate]);

  if (error) {
    return (
      <div className="min-h-[80vh] flex items-center justify-center px-4">
        <div className="card-yellow max-w-sm w-full text-center py-12 px-8">
          <div className="w-24 h-24 mx-auto mb-6 rounded-full border-4 border-red-500 flex items-center justify-center">
            <XCircle className="h-12 w-12 text-red-500" />
          </div>
          <h2 className="font-handwritten text-2xl font-bold text-red-400 mb-2">Oops</h2>
          <p className="text-gray-500 font-handwritten text-base mb-8">
            something went wrong,<br />click try again.
          </p>
          <button
            onClick={() => { window.location.href = '/api/patient/login'; }}
            className="w-full bg-red-600 text-white py-3.5 rounded-full font-handwritten text-lg hover:bg-red-700 transition-colors"
          >
            Try Again
          </button>
        </div>
      </div>
    );
  }

  return <LoadingSpinner message="Authenticating with Google..." />;
}
