if __FILE__ == $0
    filename = ARGV[0]
    colname = ARGV[1]
    InsertDataService.insert_csv!(filename, colname)
    pp DataRow.count
end
