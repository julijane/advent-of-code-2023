package main

import (
	"math"
	"slices"

	"github.com/julijane/advent-of-code-2023/aoc"
)

type Section struct {
	From, To int
}

type RangeDelta struct {
	From, To int
	Delta    int
}

type (
	Block             []*RangeDelta
	InstructionBlocks [7]*Block
)

// genInstructionsBlock converts a list of instructions from the input
// into a sorted list of ranges with the delta for this range
func genInstructionBlock(instructionLines []string) *Block {
	var block Block

	// Convert Input to Ranges
	for _, instructionLine := range instructionLines {
		instructionData := aoc.ExtractNumbers(instructionLine)

		block = append(block, &RangeDelta{
			From:  instructionData[1],
			To:    instructionData[1] + instructionData[2] - 1,
			Delta: instructionData[0] - instructionData[1],
		})
	}

	// Sort Ranges
	slices.SortFunc(block, func(i, j *RangeDelta) int {
		if i.From < j.From {
			return -1
		}
		return 1
	})

	return &block
}

// transformMappingTable applies one instruction block to a mapping table to
// get a new mapping table (i.e. a table that maps input ranges to delta values
// which incorporate the instructions from this block
func transformMappingTable(mappingTable *Block, instructionBlock *Block) *Block {
	newTable := Block{}

	for _, mappingEntry := range *mappingTable {

		// remainder saves how much of the current mappingEntry we have not yet processed
		// we add the current delta sum so that the range is correct for the input to
		// this stage. It will be subtracted later again.
		remainder := &Section{
			From: mappingEntry.From + mappingEntry.Delta,
			To:   mappingEntry.To + mappingEntry.Delta,
		}

		// Now apply instructions until we have processed the whole mappingEntry
		for _, instruction := range *instructionBlock {
			if instruction.To < remainder.From {
				// instruction range ends before this mapping entry starts, so skip
				continue
			}

			if instruction.From > remainder.To {
				// instruction range starts after this mapping entry ends, so we are done
				break
			}

			if instruction.From > remainder.From {
				// The beginning of this mapping entry stays at the same delta
				newTable = append(newTable, &RangeDelta{
					From:  remainder.From - mappingEntry.Delta,
					To:    instruction.From - mappingEntry.Delta - 1,
					Delta: mappingEntry.Delta,
				})
				remainder.From = instruction.From
			}

			// now apply the additional delta from the instruction
			newTable = append(newTable, &RangeDelta{
				From:  remainder.From - mappingEntry.Delta,
				To:    min(instruction.To, remainder.To) - mappingEntry.Delta,
				Delta: mappingEntry.Delta + instruction.Delta,
			})

			// if we end after the end of the instruction range, we leave the rest for the
			// next instruction
			if instruction.To < remainder.To {
				remainder.From = instruction.To + 1
			} else {
				// otherwise we have no remainder
				remainder = nil
				break
			}
		}

		// If we still have a remainder, append it with unmodified delta
		if remainder != nil {
			newTable = append(newTable, &RangeDelta{
				From:  remainder.From - mappingEntry.Delta,
				To:    remainder.To - mappingEntry.Delta,
				Delta: mappingEntry.Delta,
			})
		}

	}

	return &newTable
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	resultPart1 := math.MaxInt
	resultPart2 := math.MaxInt

	// Convert Instructions
	textBlocks := input.TextBlocks()
	instructionBlocks := InstructionBlocks{}
	for block := 1; block < 8; block++ {
		instructionBlocks[block-1] = genInstructionBlock(textBlocks[block][1:])
	}

	// The Mapping Table maps input ranges to delta values
	// (which are to be added in this range) to get from input to output.
	// We start with a single entry convering the full integer range,
	// which will be split up by applying the instruction blocks.
	mappingTable := &Block{
		&RangeDelta{From: 0, To: math.MaxInt, Delta: 0},
	}

	// Now apply instruction blocks to get final input to output mapping
	for _, instructionBlock := range instructionBlocks {
		mappingTable = transformMappingTable(mappingTable, instructionBlock)
	}

	// part 1
	// this could also be solved by the optimized method of part 2, but after all
	// the hours of writing the optimized code I want to keep the brute force solution
	// at least for this :) This is still somewhat optimized because we use the
	// computed mapping table.
	seeds := aoc.ExtractNumbers(textBlocks[0][0])
	for _, seed := range seeds {
		for _, mappingEntry := range *mappingTable {
			if seed >= mappingEntry.From && seed <= mappingEntry.To {
				seed += mappingEntry.Delta
				break
			}
		}

		if seed < resultPart1 {
			resultPart1 = seed
		}
	}

	// part2

	// sort the mapping table so that the mappings that yield the lowest results are first
	slices.SortFunc(*mappingTable, func(i, j *RangeDelta) int {
		if i.From+i.Delta < j.From+j.Delta {
			return -1
		}
		return 1
	})

	// Go through the mapping table until we find one that applies to one of our inputs,
	// then we have the result
search:
	for _, mappingEntry := range *mappingTable {
		for num := 0; num < len(seeds); num += 2 {
			firstSeed := seeds[num]
			numSeeds := seeds[num+1]

			// this mapping entry does not apply to this seed range
			if firstSeed > mappingEntry.To || firstSeed+numSeeds < mappingEntry.From {
				continue
			}

			// if our seed range starts before the start of the mapping entry
			// we use the start of the mapping entry to get the location
			// otherwise we get the location from the first seed of the range
			resultPart2 = max(firstSeed, mappingEntry.From) + mappingEntry.Delta
			break search
		}
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
