require_relative "binary_search_tree"

describe BinaryTree do
  context "with a depth of 3" do
    let(:n5) { Node.new(7) }
    let(:tree) do
      n4 = Node.new(15)
      n3 = Node.new(20, n4, n5)
      n2 = Node.new(9)
      root = Node.new(3, n2, n3)
      BinaryTree.new(root)
    end

    describe "depth" do
      it "is correct" do
        expect(tree.depth).to eq(3)
      end
    end
  end

  context "with a depth of 4" do
    let(:tree) do
      n6 = Node.new(9)
      n5 = Node.new(7, n6)
      n4 = Node.new(15)
      n3 = Node.new(20, n4, n5)
      n2 = Node.new(9)
      root = Node.new(3, n2, n3)
      BinaryTree.new(root)
    end

    describe "depth" do
      it "is correct" do
        expect(tree.depth).to eq(4)
      end
    end
  end
end

describe BinarySearchTree do
  context "with a valid BST" do
    let(:root) { Node.new(3, nil, nil) }
    let(:tree) do
      described_class.new(root)
    end

    describe "search" do
      it "is correct" do
        expect(tree.search(3)).to eq(root)
      end
    end

    describe "valid?" do
      it "is true" do
        expect(tree.valid?(root)).to be true
      end
    end
  end

  context "with an invalid BST" do
    let(:n5) { Node.new(7) }
    let(:n4) { Node.new(15) }
    let(:n3) { Node.new(20, n4, n5) }
    let(:n2) { Node.new(9) }
    let(:root) { Node.new(3, n2, n3) }
    let(:tree) do
      described_class.new(root)
    end

    describe "depth" do
      it "is correct" do
        expect(tree.depth).to eq(3)
      end
    end

    describe "invalid_nodes" do
      it "is correct" do
        expect(tree.invalid_nodes(root)).to eq([9, 7])
      end
    end

    describe "valid?" do
      it "is false" do
        expect(tree.valid?(root)).to be false
      end
    end
  end
end
