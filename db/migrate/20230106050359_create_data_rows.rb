class CreateDataRows < ActiveRecord::Migration[7.0]
  def change
    create_table :data_rows do |t|
      t.text :text

      t.timestamps
    end
  end
end
