require_relative './spec_helper'
require_relative '../lib/card'

describe Card do

  subject { Card.new('A', 'S') }

  it 'has a face' do
    subject.face.must_equal('A')
  end

  it 'has a suit' do
    subject.suit.must_equal('S')
  end

  it 'has a short name' do
    subject.to_s.must_equal('AS')
  end

  it 'has a long name' do
    subject.to_s_long.must_equal("ace\nspades")
  end
end
