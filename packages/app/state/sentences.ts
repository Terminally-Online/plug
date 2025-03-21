import { atom } from 'jotai'

// Store validation state for each sentence
export type SentenceValidState = {
  isValid: boolean
  isComplete: boolean
  chains: string
  actionPreview: string
}

// Key format: "${plugId}-${actionIndex}"
export const sentenceValidStateAtom = atom<Record<string, SentenceValidState>>({})

// Selector to get validation state for all sentences of a specific plug
export const plugSentencesValidStateAtom = atom(
  (get) => {
    const sentences = get(sentenceValidStateAtom)
    return (plugId: string) => {
      return Object.entries(sentences)
        .filter(([key]) => key.startsWith(`${plugId}-`))
        .map(([, value]) => value)
    }
  }
)

// Selector to check if all sentences of a specific plug are valid
export const areAllSentencesValidAtom = atom(
  (get) => {
    const getSentences = get(plugSentencesValidStateAtom)
    return (plugId: string) => {
      const sentences = getSentences(plugId)
      return sentences.length > 0 && sentences.every(s => s.isValid && s.isComplete)
    }
  }
)