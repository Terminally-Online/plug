import { useState } from 'react';
import RopeEditor from './components/RopeEditor';
import { RopeProvider } from './context/RopeContext';

function App() {
  const [templateSentences, setTemplateSentences] = useState<string[]>([
    'Swap {0} for {1} on {2}',
    'Provide liquidity for {0} on {1}',
    'Transfer {0} to {1}'
  ]);

  const addNewSentence = () => {
    setTemplateSentences([...templateSentences, 'New action with {0} and {1}']);
  };

  return (
    <RopeProvider initialSentences={templateSentences}>
      <div className="flex h-screen w-full overflow-hidden">
        <RopeEditor />
      </div>
    </RopeProvider>
  );
}

export default App;