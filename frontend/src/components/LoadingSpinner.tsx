import { Loader2 } from 'lucide-react';

export default function LoadingSpinner({ message = 'loading...' }: { message?: string }) {
  return (
    <div className="flex flex-col items-center justify-center py-12">
      <Loader2 className="h-8 w-8 text-brand-black animate-spin" />
      <p className="mt-3 text-sm text-gray-500 font-handwritten">{message}</p>
    </div>
  );
}
