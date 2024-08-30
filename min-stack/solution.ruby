class MinStack
  def initialize()
      @stacks = []
      @minStackVal = []
  end

  def push(val)
      if @stacks.length == 0 
          @minStackVal << val
      else 
          @minStackVal << [@minStackVal[@stacks.length - 1], val].min
      end

      @stacks << val
  end

  def pop
   if @stacks.length > 0 
      @stacks.delete_at(@stacks.length - 1)
      @minStackVal.delete_at(@stacks.length)
   end
  end
 
 def top
      @stacks[@stacks.length - 1]
 end

 def get_min()
      @minStackVal[@stacks.length - 1]
 end
end
