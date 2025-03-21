import RopeEditor from './components/RopeEditor';
import { RopeProvider } from './context/RopeContext';

function App() {
  return (
    <RopeProvider>
      <div className="flex h-screen w-full overflow-hidden">
        <RopeEditor />
      </div>
    </RopeProvider>
  );
}

export default App;