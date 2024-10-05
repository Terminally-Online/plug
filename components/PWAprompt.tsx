import React, { useState, useEffect } from 'react';
import { isWebIOS, isWebAndroid } from '@/lib/utils/platform';

// Add this type declaration at the top of your file or in a separate .d.ts file
declare global {
  interface WindowEventMap {
    beforeinstallprompt: Event;
  }
}

const PWAPrompt: React.FC = () => {
  const [showPrompt, setShowPrompt] = useState(false);
  const [deferredPrompt, setDeferredPrompt] = useState<any>(null);

  useEffect(() => {
    const handleBeforeInstallPrompt = (e: Event) => {
      e.preventDefault();
      setDeferredPrompt(e);
      setShowPrompt(true);
    };

    window.addEventListener('beforeinstallprompt', handleBeforeInstallPrompt as EventListener);

    return () => {
      window.removeEventListener('beforeinstallprompt', handleBeforeInstallPrompt);
    };
  }, []);

  const handleInstall = () => {
    if (deferredPrompt) {
      deferredPrompt.prompt();
      deferredPrompt.userChoice.then((choiceResult: { outcome: string }) => {
        if (choiceResult.outcome === 'accepted') {
          console.log('User accepted the install prompt');
        } else {
          console.log('User dismissed the install prompt');
        }
        setDeferredPrompt(null);
      });
    }
  };

  if (!showPrompt) return null;

  if (isWebIOS) {
    return (
      <div>
        <p>To install the app, tap the share button and then "Add to Home Screen"</p>
      </div>
    );
  }

  if (isWebAndroid) {
    return (
      <div>
        <p>Install Plug for a better experience</p>
        <button onClick={handleInstall}>Install</button>
      </div>
    );
  }

  return null;
};

export default PWAPrompt;
