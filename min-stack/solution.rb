class MinStack
  def initialize
    @stacks = []
    @minStackVal = []
  end

  def push(val)
    @minStackVal << if @stacks.length == 0
                      val
                    else
                      [@minStackVal[@stacks.length - 1], val].min
                    end

    @stacks << val
  end

  def pop
    return unless @stacks.length > 0

    @stacks.delete_at(@stacks.length - 1)
    @minStackVal.delete_at(@stacks.length)
  end

  def top
    @stacks[@stacks.length - 1]
  end

  def get_min
    @minStackVal[@stacks.length - 1]
  end
end
