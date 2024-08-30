#rubocop:disable all
class Trie
    def initialize()
        @root = Node.new(nil)
    end

    def insert(word)
      curr = @root
      i = 0

      while i < word.length 
        if !curr.children[word[0..i]] 
         curr.children[word[0..i]] = Node.new(word[0..i])
        end
        curr = curr.children[word[0..i]]
        i+=1
      end

      curr.is_final = true
    end

    def search(word)
      curr = @root.children
      i = 0

      while i < word.length - 1
        if !curr[word[0..i]]
          return false
        end

        curr = curr[word[0..i]].children

        i += 1
      end

      return curr[word].is_final
    end

    def startsWith(word)
      curr = @root.children
      i = 0

      while i < word.length - 1
        if !curr[word[0..i]]
          return false
        end

        curr = curr[word[0..i]].children

        i += 1
      end

      return curr[word] != nil
    end
end

class Node
 attr_accessor :val, :is_final, :children
 def initialize(val, is_final = false)
  @val = val
  @is_final = is_final
  @children = {}
 end 
end

trie = Trie.new 

trie.insert("foo")
trie.insert("bar")
trie.insert("foego")

pp trie.search("foo")
pp trie.search("bar")
pp trie.search("foego")
pp trie.search("lucas")
pp trie.search("ba")

pp trie.startsWith("foego")
