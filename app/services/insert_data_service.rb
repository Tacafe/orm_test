require 'csv'

class InsertDataService
  def self.insert_csv!(filepath, colname)
    CSV.open(filepath, headers: true).each do |row|
      ActiveRecord::Base.transaction do
        DataRow.create!(text: row[colname])
      end
    end
  end
end
