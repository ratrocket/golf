class Card
  attr_reader :value, :suit, :face

  @@values = {
    'A' => 1, '2' => 2,  '3' => 3,  '4' => 4,
    '5' => 5, '6' => 6,  '7' => 7,  '8' => 8,
    '9' => 9, 'T' => 10, 'J' => 10, 'Q' => 10,
    'K' => 0, '0' => -5
  }

  @@long_face_names = {
    'A' => 'ace', '2' => 'two', '3' => 'three', '4' => 'four',
    '5' => 'five', '6' => 'six', '7' => 'seven', '8' => 'eight',
    '9' => 'nine', 'T' => 'ten', 'J' => 'jack', 'Q' => 'queen',
    'K' => 'king', '0' => 'joker'
  }

  @@long_suit_names = {
    'H' => 'hearts', 'D' => 'diamonds', 'C' => 'clubs',
    'S' => 'spades', 'X' => 'none'
  }

  # f, s: strings (see DATA portion for allowable values)
  def initialize(f, s)
    if self.valid?(f, s)
      @face = f
      @suit = s
      @value = @@values[@face]
    else
      raise "Invalid card"
    end
  end

  def to_s
    "#{@face}#{@suit}"
  end

  def to_s_long
    "#{@@long_face_names[@face]}\n#{@@long_suit_names[@suit]}"
  end

  private

  def self.valid?(f, s)
    valid_combination(f, s) && valid_faces.include?(f) && valid_suits.include?(s)
  end

  def self.invalid_combination?(f, s)
    s == 'X' && f != '0'  # suit is 'none' & face is not joker
  end

  def self.valid_combination?(f, s)
    !invalid_combination?
  end

  def valid_suits
    @valid_suits ||= @@long_suit_names.keys
  end

  def valid_faces
    @valid_faces ||= @@long_face_names.keys
  end
end

__END__

faces:
  %w(A 2 3 4 5 6 7 8 9 T J Q K 0)
where T is 10
      0 is joker

suits:
  %w(H D C S X)  # hearts, diamonds, clubs, spades, X for no suit (jokers)

values:
A 1
2 2
3 3
4 4
5 5
6 6
7 7
8 8
9 9
T 10
J 10
Q 10
K 0
0 -5
