import { Link } from 'react-router-dom';

export default function Landing() {
  return (
    <div className="min-h-screen bg-brand-yellow-pale">
      {/* Header */}
      <div className="text-center pt-10 pb-4">
        <p className="text-sm text-gray-500 font-sans mb-1">Get in touch</p>
        <h1 className="brand-name text-5xl">LifeLink</h1>
      </div>

      {/* Tagline */}
      <div className="max-w-xl mx-auto text-center px-4 mb-12">
        <p className="text-gray-600 text-sm font-sans leading-relaxed">
          Your confidential gateway to professional mental health support, providing
          seamless video streaming and real-time chat with experienced psychologists, all
          from the comfort of your own space.
        </p>
      </div>

      {/* Who's Here */}
      <div className="text-center mb-12">
        <h2 className="font-handwritten text-4xl font-bold tracking-wide mb-8 uppercase">
          Who's Here?
        </h2>

        <div className="flex justify-center gap-12 mb-10">
          {/* Patient */}
          <div className="flex flex-col items-center gap-3">
            <div className="w-28 h-28 rounded-full bg-brand-yellow border-4 border-brand-yellow flex items-center justify-center text-5xl shadow-sm hover:scale-105 transition-transform cursor-pointer">
              <span role="img" aria-label="patient">&#x1F60A;</span>
            </div>
            <a href="/api/patient/login" className="btn-dark text-sm py-2 px-6">
              Patient
            </a>
          </div>

          {/* Doctor */}
          <div className="flex flex-col items-center gap-3">
            <div className="w-28 h-28 rounded-full bg-orange-300 border-4 border-white flex items-center justify-center text-5xl shadow-sm hover:scale-105 transition-transform cursor-pointer">
              <span role="img" aria-label="doctor">&#x1F60A;</span>
            </div>
            <Link to="/doctor/login" className="btn-dark text-sm py-2 px-6">
              Doctor
            </Link>
          </div>
        </div>
      </div>

      {/* Support Center */}
      <div className="text-center mb-10">
        <Link to="/admin/login" className="btn-outline text-sm">
          Support Center
        </Link>
      </div>

      {/* Emoji Image Section */}
      <div className="max-w-md mx-auto px-4 mb-8">
        <div className="rounded-2xl overflow-hidden shadow-lg">
          <div className="bg-brand-black flex items-center justify-center py-12 gap-6">
            <div className="text-6xl opacity-90">&#x1F61E;</div>
            <div className="text-6xl opacity-90">&#x1F610;</div>
            <div className="text-6xl opacity-90">&#x1F60A;</div>
          </div>
        </div>
      </div>

      {/* Send Message */}
      <div className="text-center pb-16">
        <button className="btn-dark text-sm">
          Send message
        </button>
      </div>
    </div>
  );
}
