=begin
20210708にゼミでやった。
母数に対して世論調査でどのくらいの人数調査すべきか出す。
https://fukusima-zemi.vercel.app/finite_population にweb上でできるように実装
=end

class FinitePopulation
  def output
    mothers / children
  end

  def mothers
    16000
  end

  def child1
    0.05 / 1.96
  end

  def child2
    mother = mothers - 1
    mother / 0.25
  end

  def children
    child1 * child1 * child2 + 1
  end
end

p FinitePopulation.new.output