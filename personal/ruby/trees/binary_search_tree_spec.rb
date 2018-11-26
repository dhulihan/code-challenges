require_relative "binary_search_tree"

describe BinaryTree do
  describe "parse_flat_array" do
    let(:input) { [3, 6, 2, 9, -1, 10] }

    it "works as expected" do
      root = described_class.parse_flat_array(input)
      p root
      expect(root.depth).to eq(3)
    end

  end

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

    describe
  end
end

describe BinarySearchTree do
  context "with a valid BST" do
    let(:n6) { Node.new(6, nil, nil) }
    let(:n5) { Node.new(7, n6, nil) }
    let(:n4) { Node.new(4, nil, nil) }
    let(:n3) { Node.new(5, n4, n5) }
    let(:n2) { Node.new(2, nil, nil) }
    let(:root) { Node.new(3, n2, n3) }
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

    describe "remove" do
      let(:tree_after) do
        n5 = Node.new(7, nil, nil)
        n4 = Node.new(4, nil, nil)
        n6 = Node.new(6, n4, n5)
        n2 = Node.new(2, nil, nil)
        root = Node.new(3, n2, n6)
        described_class.new(root)
      end

      it "transforms the tree as expected" do
        tree.delete(n3.value)
        expect(tree).to eq tree_after
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
