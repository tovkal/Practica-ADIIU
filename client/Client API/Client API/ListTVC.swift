//
//  ListTVC.swift
//  Client API
//
//  Created by Andrés Pizá on 20/1/15.
//  Copyright (c) 2015 tovkal. All rights reserved.
//

import UIKit
import SwiftyJSON
import Alamofire

class ListTVC: UITableViewController, UITableViewDataSource {
    
    var operationTitle: String?
    var method: String?
    
    var tableData: JSON = JSON.nullJSON

    @IBOutlet weak var list: UITableView!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.title = operationTitle
        
        fetchData()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }

    // MARK: - Table view data source

    override func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        return 1
    }

    override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return tableData.count
    }

    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCellWithIdentifier("listCell", forIndexPath: indexPath) as UITableViewCell
        
        cell.textLabel?.text = self.tableData[indexPath.row]["nombre"].string!
        cell.detailTextLabel?.text = self.tableData[indexPath.row]["texto"].string!

        return cell
    }
    
    // MARK: - API Client
    
    func fetchData() {
        var url: String = "http://localhost:8080/api" + method!
        
        Alamofire.request(.GET, url, parameters: nil)
            .responseJSON { (req, res, json, error) in
                if(error != nil) {
                    NSLog("Error: \(error)\nfor url: \(url)")
                }
                else {
                    NSLog("Success: \(url)")
                    self.tableData = JSON(json!)
                    self.list.reloadData()
                }
        }
    }
}
