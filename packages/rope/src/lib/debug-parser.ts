import { SentenceService } from '../services/sentence';

// Create an instance of the service
const sentenceService = new SentenceService();

// The problematic sentence
const sentence = 'Get balance of {0<token:address:uint256:uint256>} held by {1<holder:string>}';

try {
  // Try to parse the sentence
  const parsed = sentenceService.parseSentence(sentence);
  
  console.log('Successfully parsed sentence:');
  console.log('Original:', parsed.original);
  console.log('Template:', parsed.template);
  console.log('Parts:', parsed.parts);
  console.log('Inputs:', JSON.stringify(parsed.inputs, null, 2));
} catch (error) {
  console.error('Error parsing sentence:', error);
}

// Let's also test with a simpler sentence
const simpleSentence = 'Swap {0<amount:uint256>} {1<token:address>}';

try {
  // Try to parse the simpler sentence
  const parsed = sentenceService.parseSentence(simpleSentence);
  
  console.log('\nSuccessfully parsed simple sentence:');
  console.log('Original:', parsed.original);
  console.log('Template:', parsed.template);
  console.log('Parts:', parsed.parts);
  console.log('Inputs:', JSON.stringify(parsed.inputs, null, 2));
} catch (error) {
  console.error('Error parsing simple sentence:', error);
}